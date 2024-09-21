import { createRouter, createWebHistory } from 'vue-router'
import Body from '@/components/Body.vue'
import LoginComposition from '@/components/LoginComposition.vue'
import BooksComposition from '@/components/BooksComposition.vue'
import BookComposition from '@/components/BookComposition.vue'
import BooksAdmin from '@/components/BooksAdmin.vue'
import BookEdit from '@/components/BookEdit.vue'
import Users from '@/components/Users.vue'
import User from '@/components/UserEdit.vue'
import Security from '@/components/security.js'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Body
  },
  {
    path: '/login',
    name: 'LoginComposition',
    component: LoginComposition
  },
  {
    path: '/books',
    name: 'BooksComposition',
    component: BooksComposition
  },
  {
    path: '/books/:bookName',
    name: 'BookComposition',
    component: BookComposition
  },
  {
    path: '/admin/books',
    name: 'BooksAdmin',
    component: BooksAdmin
  },
  {
    path: '/admin/books/:bookId',
    name: 'BookEdit',
    component: BookEdit
  },
  {
    path: '/admin/users',
    name: 'Users',
    component: Users
  },
  {
    path: '/admin/users/:userId',
    name: 'User',
    component: User
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// this function will run every time a user navigates to any part of the site,
// and call checkToken(). If the user has a token in store.token, then we validate
// that token.
router.beforeEach(() => {
  Security.checkToken()
})

export default router
