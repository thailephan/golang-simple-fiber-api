# golang-simple-fiber-api
Simple web api with fiber package, and mongodb

## Project structure
1. **repository**: Contains code for connect, crud in db
2. **repository/connect.go**: Connect database, and assign global `Client`, `TodoCollection` and `Ctx`
3. **repository/repository.go**: Want to define default interface for all repo but don't know how
4. **repository/flashcard_repository.go**: Contains functions to CRUD to database
5. **mocks**: Contains mock datas for project
6. **models**: Contains models/entities for project
7. **usecases**: Contains usecases (actions) of models/entities
8. **usecases/flashcard_usecase.go**: Contains usecases (actions) of models/entities for flashcard
9. **pkg**: Don't know
10. **pkg/utils**: Utils for project
10. **pkg/utils/parse.go**: Parse from string to int function
10. **pkg/utils/primitive.go**: Utils for mongodb modules: `ObjectIDFromHex`, `NewObjectID`, `BsonMapToEntity`
10. **pkg/utils/read_json.go**: Utils for read json file from path
10. **pkg/utils/query_options**: Utils for `GetAll` query like `limit`, `offset`
10. **pkg/utils/json_response.go**: Utils for json response template
8. **docker-compose.yml**: Define mongodb service to run

## Author
Le Phan Minh Thai

## License
MIT