curl -XPOST -H "Content-Type: application/json" --data '{
    "query": "mutation AddUser($name: String, $age: Int) {addUser(name: $name, age: $age)}",
    "variables": {
        "name": "test",
        "age": 33
    }
}' http://localhost:8080/graphql