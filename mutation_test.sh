curl -XPOST -H "Content-Type: application/json" --data '{
    "query": "mutation AddItem($title: String, $price: Float) {addItem(title: $title, price: $price)}",
    "variables": {
        "title": "New Item",
        "price": 10000
    }
}' http://localhost:8080/graphql