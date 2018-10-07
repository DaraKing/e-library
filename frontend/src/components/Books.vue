<template>
    <div>
        <div class="title">
            <div class="titleContent">All books</div>
            <router-link to="/add-book">
                <div class="titleContent"><img src="../assets/plus.svg" class="icon"/></div>
            </router-link>
        </div>
        <div class="book-list">
            <div class="book" v-for="book in books">
                <router-link :to="{name: 'singleBook', params: {bookId: book.id} }">
                    <div class="book-img">
                        <img :src="'http://localhost:3030/images/'+book.image_name" class="book-image">
                    </div>
                    <div class="book-info">
                        {{ book.title }}
                    </div>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Books",
        data() {
            return {
                books: [],
            }
        },
        created: function() {
            this.getBooks();
        },
        methods: {
            getBooks: function () {
                this.$http.get('http://localhost:3030/get-books').then(response => {
                    this.books = response.body;
                    console.log(this.books);
                })
            }
        }
    }
</script>

<style scoped>
    @import url('https://fonts.googleapis.com/css?family=Noto+Serif+KR');

    .title {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    .titleContent {
        padding: 25px;
    }

    .icon {
        height: 36px;
    }

    .book-list {
        display: flex;
        flex-direction: row;
        height: auto;
    }

    .book {
        background-size: 100% 100%;
        width: 25%;
        height: 400px;
    }

    .book-img {
        height: 95%;
    }

    .book-image {
        width: 100%;
        height: 100%;
    }

    .book-info {
        font-family: 'Noto Serif KR', sans-serif;
        font-size: 24px;
        color: #fff;
    }
</style>