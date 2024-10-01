# User Usecase FindOne

```
FindOne(ctx context.Context, id int) (GetOneResp, apperr.Err)
```

### Response
```
type GetOneResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

```mermaid
sequenceDiagram
  actor c As Caller
  participant u AS UserUsecase
  participant ur AS UserRepo

  Note over c: GET /users/:id
  c ->> u: ctx, id
  Note left of u: FindOne

  u ->> ur: ctx, id
  Note left of ur: FindOne
  ur -->> u: user, nil
  note left of ur: FindOne
  u -->> c: {id, name, email}, nil
  note over c: 200

  break user not found
    ur -->> u: nil, nil
    note left of ur: FindOne
    u -->> c: {}, err<br/>User Not found<br/>ไม่พบ User
    note left of u: 404
  end

  break generic db error
    ur -->> u: {}, err<br/> Error
    note left of ur: FindOne
    u -->> c: {}, err<br/> Error
    note left of u: 500
  end
```