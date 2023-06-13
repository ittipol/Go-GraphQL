curl --location --request POST 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query": "query DATA($slug: String) {getItemBySlug(slug: $slug) {id title} allCategories{id name}}",
    "variables": {
        "slug": "44xzc84mRSk"
    }
}'