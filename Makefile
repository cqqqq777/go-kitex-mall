# run the user
.PHONY: user
user:
	go run ./cmd/user

# run the merchant
.PHONY: merchant
merchant:
	go run ./cmd/merchant


# run the product
.PHONY: product
product:
	go run ./cmd/product


# run the operate
.PHONY: operate
operate:
	go run ./cmd/operate


# run the order
.PHONY: order
order:
	go run ./cmd/order


# run the pay
.PHONY: pay
pay:
	go run ./cmd/pay