# calorie-tracker-Go-react


Calorie tracker is Go backend for adding the calories of food.


## APIs

```
GET http://localhost:8000/entries
```



```
POST http://localhost:8000/entry/create

body = 
{
    "dish": "chapathi",
    "fat": 5,
    "ingredients": "sugar",
    "calories": "30"
}
```


```
PUT http://localhost:8000/incredient/update/:id
 body = 
{
    "ingredients": "salt"
}
```

```
DELETE http://localhost:8000/entry/delete/:id
```
Other APIs are updating entries
