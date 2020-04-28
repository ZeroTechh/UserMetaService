
# UserMetaService

A velocity which will be handling user meta data

## Functions

### Add

Creates meta data for a user and adds it into database.

```go
_, err := client.Add(ctx, &proto.Identifier{UserID: "user id"})
```

### Get

Finds and returns meta data of a user from database.

```go
metaData, err := client.Get(ctx, &proto.Identifier{UserID: "user id"})
```

### Activate

Sets user status as verified.

```go
_, err := client.Activate(ctx, &proto.Identifier{UserID: "user id"})
```

### Delete

Sets user status as deleted.

```go
_, err := client.Delete(ctx, &proto.Identifier{UserID: "user id"})
```
