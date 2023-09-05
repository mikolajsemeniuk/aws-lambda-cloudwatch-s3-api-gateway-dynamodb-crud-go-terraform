locals {
  routes = [
    {
      name   = "doc",
      method = "GET",
      route  = "/doc",
    },
    {
      name   = "list",
      method = "GET",
      route  = "/orders",
    },
    {
      name   = "create",
      method = "POST",
      route  = "/orders",
    },
    {
      name   = "find",
      method = "GET",
      route  = "/orders/{id}",
    },
    {
      name   = "update",
      method = "PUT",
      route  = "/orders/{id}",
    },
    {
      name   = "remove",
      method = "DELETE",
      route  = "/orders/{id}",
    },
  ]
}
