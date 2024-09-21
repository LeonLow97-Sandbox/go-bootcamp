# Course Content

- Working with Vue 3 using a CDN
- Working with Vue 3 using the vue-cli, node.js and npm
- Learn both the Options API and the Composition API for Vue 3
- How to work with props
- How to build reusable Vue components
- How to build and use a data store with Vue 3
- Creating, validating, and posting forms using fetch and JSON
- Emitting and processing events in Vue
- Conditional rendering in Vue
- Animations and transitions in Vue
- Working with the Vue Router
- Protecting routes in Vue (requiring user authentication)
- Caching components using Vue's KeepAlive functionality
- Implementing a REST API using Go
- Routing with Go
- Connecting to a Postgresql database with Go
- Reading and writing JSON with Go
- Complete user authentication with Go using stateful tokens
- Testing our Go back end with unit and integration tests

# Vue

## Making **JWT Token** a global variable

```js
// store.js to store the value of the token
import { reactive } from "vue";

export const store = reactive({
  token: "",
});
```

## Navigation Guards

- Navigation guards provided by Vue router are primarily used to guard navigation either by redirecting it or canceling it.
- Protect the routes when user is unauthorized.

```js
const router = createRouter({ ... })

router.beforeEach((to, from) => {
  // ...
  // explicitly return false to cancel the navigation
  return false
})
```

```js
//App.vue
<router-view
  v-slot="{ Component }"
  :key="componentKey"
  @success="success"
  @error="error"
  @warning="warning"
  @forceUpdate="forceUpdate"
>
  <keep-alive>
    <component :is="Component" />
  </keep-alive>
</router-view>
```

## `TransitionGroup`

- [TransitionGroup](https://vuejs.org/guide/built-ins/transition-group.html)
- [Transition](https://vuejs.org/guide/built-ins/transition.html)

## `KeepAlive`

- `<KeepAlive>` is a built-in component that allows us to conditionally cache component instances when dynamically switching between multiple components.
- [KeepAlive](https://vuejs.org/guide/built-ins/keep-alive.html)

```js
<KeepAlive>
  <component :is="activeComponent" />
</KeepAlive>
```

- `activated` lifecycle hook is called when a component inside a `keep-alive` becomes the active component to be displayed.
- `deactivated` lifecycle hook is called when a component inside a `keep-alive` is no longer the active component and is being cached for potential reuse.

```js
// Book.vue
<script setup>
export default {
  activated() {
    // Component has been re-activated
    // You can reset data or fetch new data here
  },
  deactivated() {
    // Component is being deactivated
    // You can reset data or perform cleanup here
  }
}
</script>
```

## `loadCoverImage` function (`BookEdit.vue`)

- Responsible for loading and encoding an image file selected by the user using an `input type="file">` element.
- It then converts the image into a base64 encoded string, which can be used to display or send the image data to a backend API.

```js
loadCoverImage() {
  // get a reference to the input using ref
  const file = this.$refs.coverInput.files[0]

  // encode the file using the FileReader API
  // need base64 string to send in JSON to backend API
  const reader = new FileReader()
  reader.onloadend = () => {
    const base64String = reader.result.replace('data:', '').replace(/^.+,/, '')
    this.book.cover = base64String
    // alert(base64String);
  }
  reader.readAsDataURL(file)
},
```

#### `KeepAlive` with `include`

- By default, `KeepAlive` will cache any component instance inside.
- Can customize this behavior via the `include` and `exclude` props. Both props can be a comma-delimited string, a `RegExp`, or any array containing either types.
- The match is checked against the component's `name` option, so components that need to be conditionally cached by `KeepAlive` must explicitly declare a `name` option.

```js
// Example
<KeepAlive include="a,b">
  <component :is="view" />
</KeepAlive>
```

```js
// App.vue
<keep-alive include="Books">
  <component :is="Component" />
</keep-alive>
```

## Building Application for Production

- `npm run build`
- Created a new folder `dist` in Vue root directory

```js
// package.json
"scripts": {
  "dev": "vite",
  "build": "vite build --mode production",
}
```

- We used `Caddy` for our web server.
- [Other Web Servers with Vue](https://router.vuejs.org/guide/essentials/history-mode.html#example-server-configurations)

# Golang

## How Go works with JSON Files

---

#### Useful Links

- [Go Playground MarshalIndent](https://go.dev/play/p/2SP7H3ybk9D)
- [Go Playground Unmarshal](https://go.dev/play/p/NaMNv6bzEXA)
- [GoByExample JSON](https://gobyexample.com/json)
- [Go Default Encoding Types (Go Type - JSON Type)](https://blog.boot.dev/golang/json-golang/#default-type)

---

#### Go Default Encoding Types

JSON and GO Types don't match up 1 to 1. Below is a table that describes the type relationships when encoding and decoding.

|   Go Type   | JSON Type |
| :---------: | :-------: |
|    bool     |  boolean  |
|   float64   |  number   |
|   string    |  string   |
| nil pointer |   null    |

---

## Error Codes in PostgreSQL

- [Error Codes in Postgres](https://www.postgresql.org/docs/current/errcodes-appendix.html)

## Decoding an image base64 encoded string

```go
// handlers.go
var staticPath = "./static/" // var because we can unit test this by changing the output path location

if len(requestPayload.CoverBase64) > 0 {
  // we have a cover
  decoded, err := base64.StdEncoding.DecodeString(requestPayload.CoverBase64)
  if err != nil {
    app.errorJSON(w, err)
    return
  }

  // write image to /static/covers
  if err := os.WriteFile(fmt.Sprintf("%s/covers/%s.jpg", staticPath, book.Slug), decoded, 0666); err != nil {
    app.errorJSON(w, err)
    return
  }
```
