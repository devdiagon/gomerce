package cart

import (
	"fmt"

	"github.com/devdiagon/gomerce/types"
)

func getCartItemsIds(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductId)
		}

		productIds[i] = item.ProductId
	}

	return productIds, nil
}

func (h *Handler) createOrder(products []types.Product, items []types.CartItem, userId int) (int, float64, error) {
	//Create a map to quickly access the product data
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	//Check the stock of the products
	if err := checkCartItemsStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	//Compute total price
	totalPrice := computeTotalPrice(items, productMap)

	//Reduce product Stock in the DB
	for _, item := range items {
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity

		h.producStore.UpdateProduct(product)
	}

	//Create the order
	orderId, err := h.orderStore.CreateOrder(types.Order{
		UserId:  userId,
		Total:   totalPrice,
		Status:  "pending",
		Address: "HARDCODED ADDRESS",
	})

	if err != nil {
		return 0, 0, err
	}

	//Create order items
	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductId].Price,
		})
	}

	return orderId, totalPrice, nil
}

func checkCartItemsStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductId]
		if !ok {
			return fmt.Errorf("product: %s, is not available in the store, please refresh your cart", product.Name)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("there is not enough stock for the product %s", product.Name)
		}
	}

	return nil
}

func computeTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductId]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
