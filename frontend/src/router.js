import Vue from 'vue'
import Router from 'vue-router'
import Register from './components/Register'
import Login from './components/Login'
import Logout from './components/Logout'
import Home from './components/Home'
import Books from './components/Books'
import AddBook from './components/AddBook'
import SingleBook from './components/singleBook'
import ReservedBooks from './components/ReservedBooks'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/logout',
      name: 'logout',
      component: Logout
    },
    {
      path: '/books',
      name: 'books',
      component: Books
    },
    {
      path: '/add-book',
      name: 'addBook',
      component: AddBook,
    },
    {
      path: '/singleBook/:bookId',
      name: 'singleBook',
      component: SingleBook
    },
    {
      path: '/reserved-books',
      name: 'reservedBooks',
      component: ReservedBooks
    }
  ]
})
