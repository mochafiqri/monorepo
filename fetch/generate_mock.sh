echo "==generating mockfile for repository=="
mockgen -source=./commons/interfaces/product.go -destination=./mock/product.go
mockgen -source=./commons/interfaces/currency.go -destination=./mock/currency.go
echo "==mockfile for repository generated=="
