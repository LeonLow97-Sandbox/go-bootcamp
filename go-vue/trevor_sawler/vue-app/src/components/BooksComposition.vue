<template>
  <!-- we must remove all references to `this`-->
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Books</h1>
      </div>

      <hr />
      <div class="filters text-center">
        <span class="filter me-1" :class="{ active: currentFilter === 0 }" @click="setFilter(0)"
          >ALL</span
        >
        <span class="filter me-1" :class="{ active: currentFilter === 7 }" @click="setFilter(7)"
          >CLASSIC</span
        >
        <span class="filter me-1" :class="{ active: currentFilter === 2 }" @click="setFilter(2)"
          >FANTASY</span
        >
        <span class="filter me-1" :class="{ active: currentFilter === 6 }" @click="setFilter(6)"
          >HORROR</span
        >
        <span class="filter me-1" :class="{ active: currentFilter === 4 }" @click="setFilter(4)"
          >THRILLER</span
        >
        <span class="filter me-1" :class="{ active: currentFilter === 1 }" @click="setFilter(1)"
          >SCIENCE FICTION</span
        >
      </div>
      <hr />

      <div v-if="ready">
        <div class="card-group">
          <TransitionGroup class="p-3 d-flex flex-wrap" tag="div" appear name="books">
            <div v-for="b in books" :key="b.id">
              <div
                class="card me-2 ms-1 mb-3"
                style="width: 10rem"
                v-if="b.genre_ids.includes(currentFilter) || currentFilter === 0"
              >
                <router-link :to="`/books/${b.slug}`">
                  <img
                    :src="`${imgPath}/covers/${b.slug}.jpg`"
                    class="card-img-top"
                    :alt="`cover for ${b.title}`"
                  />
                </router-link>

                <div class="card-body text-center">
                  <h6 class="card-title">{{ b.title }}</h6>
                  <span class="book-author">{{ b.author.author_name }}</span>
                  <small class="text-muted book-genre" v-for="(g, index) in b.genres" :key="g.id">
                    <em class="me-1">{{ g.genre_name }}</em>
                    <template v-if="index !== b.genres.length - 1">,</template>
                  </small>
                </div>
              </div>
            </div>
          </TransitionGroup>
        </div>
      </div>
      <p v-else>Loading...</p>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'BooksComposition',
  emits: ['error'],
  props: {},

  setup(props, ctx) {
    // setup state for this component
    let ready = ref(false)
    let currentFilter = ref(0)
    const imgPath = ref(import.meta.env.VITE_APP_IMAGE_URL)
    let books = ref({})

    // use onMounted lifecycle hook to get books
    onMounted(() => {
      console.log('using books with composition API')
      fetch(import.meta.env.VITE_APP_API_URL + '/books')
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            ctx.emit('error', response.message)
          } else {
            books.value = response.data.books
            ready.value = true
          }
        })
        .catch((error) => {
          ctx.emit('error', error)
        })
    })

    function setFilter(filter) {
      currentFilter.value = filter
    }

    // return data and functions
    return {
      ready,
      currentFilter,
      imgPath,
      books,
      setFilter
    }
  }
}
</script>

<style scoped>
.filters {
  height: 2.5em;
}
.filter {
  padding: 6px 6px;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.35s;
  border: 1px solid silver;
}

.filter.active {
  background: lightgreen;
}

.filter:hover {
  background: lightgray;
}

.book-author,
.book-genre {
  font-size: 0.8em;
}

/* transition styles in TransitionGroup */
.books-move {
  transition: all 500ms ease-in-out 50ms;
}

.books-enter-active {
  transition: all 500ms ease-in-out;
}

.books-leave-active {
  transition: all 500ms ease-in;
}

.books-enter,
.books-leave {
  opacity: 0;
}
</style>
