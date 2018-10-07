<template>
    <div>
        <div v-show="!isLoaded">
            <svg xmlns="http://www.w3.org/2000/svg" version="1.1">
                <defs>
                    <filter id="gooey">
                        <feGaussianBlur in="SourceGraphic" stdDeviation="10" result="blur"></feGaussianBlur>
                        <feColorMatrix in="blur" mode="matrix" values="1 0 0 0 0  0 1 0 0 0  0 0 1 0 0  0 0 0 18 -7" result="goo"></feColorMatrix>
                        <feBlend in="SourceGraphic" in2="goo"></feBlend>
                    </filter>
                </defs>
            </svg>
            <div class="blob blob-0"></div>
            <div class="blob blob-1"></div>
            <div class="blob blob-2"></div>
            <div class="blob blob-3"></div>
            <div class="blob blob-4"></div>
            <div class="blob blob-5"></div>
        </div>
        <div class="center-content">
            <div class="image-content" v-show="isLoaded">
                <h1>{{book.title}}</h1>
                <h4>{{book.author}}</h4>
                <img :src="'http://localhost:3030/images/'+book.image_name" class="book-image">
            </div>
            <div class="text-content" v-show="isLoaded">
                {{book.description}}

                <div class="middle-box">
                    <div class="stars-box">
                        <span class="overall-rating">{{book.rating}}</span> /<span class="max-rating">5</span>
                        <h3>Rate this book</h3>
                        <span v-for="n in 5" v-bind:id="'star-'+n" v-on:click="rateThis(n)" class="stars">
                            <font-awesome-icon :icon="['fas', 'star']" />
                        </span>
                        <div></div>
                        <div class="rate-button" v-show="isRated" v-on:click="sendRate">
                            Rate
                        </div>
                    </div>
                    <div>
                        <b>Quantity: </b> {{book.quantity}}
                        <div></div>
                        <div class="take-btn" v-if="book.quantity != 0" v-on:click="takeBook">
                            Take it
                        </div>
                    </div>
                </div>
                <div class="comment-section">
                    <div class="comment-text">
                        <input type="text" v-model="comment" class="comment-input" placeholder="Your comment">
                    </div>
                    <div class="sendBtn">
                        <img src="../assets/paper-plane.svg" class="sendImg" v-on:click="sendComment" />
                    </div>
                </div>
                <div class="comments">
                    <div class="comment" v-for="comment in book.comments">
                        {{comment.username}} - {{comment.comment_text}}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "singleBook",
        data() {
            return {
                isLoaded: false,
                book: [],
                isRated: false,
                rate: 0,
                comment: '',
            }
        },
        created: function () {
            this.getSingleBook();
        },
        methods: {
            getSingleBook: function () {
                var book_id = this.$route.params.bookId;

                this.$http.get('http://localhost:3030/book/'+book_id).then(response => {
                    this.book = response.body;
                    console.log(this.book);
                    this.isLoaded = true;
                })
            },
            rateThis: function (rate) {
                this.rate = rate;
                $(".stars").removeClass('checked');
                for(var i=1; i<=rate; i++) {
                    $("#star-"+i).addClass('checked');
                }
                this.isRated=true;
            },
            sendRate: function () {
                var rateJSON = {
                    "rate":     this.rate,
                    "bookID":   this.book.id
                };

                this.$http.post('http://localhost:3030/rate', rateJSON).then(response => {
                    if(response.status === 200) {
                        console.log('Rated');
                    }
                })
            },
            takeBook: function () {
                var takingBookJSON = {
                    "remember_token":   this.$cookie.get('rememberToken').trim().replace(/['"]+/g, ''),
                    "bookID":           this.book.id,
                };

                this.$http.post('http://localhost:3030/take-book', takingBookJSON).then(response => {
                    if(response.status === 200) {
                        console.log('Taking successful');
                    }
                })
            },
            sendComment: function() {
                var commentJSON = {
                    "remember_token":   this.$cookie.get('rememberToken').trim().replace(/['"]+/g, ''),
                    "bookID":           this.book.id,
                    "comment":          this.comment,
                };

                this.$http.post('http://localhost:3030/send-comment', commentJSON).then(response => {
                    if(response.status === 200) {
                        console.log('Successfuly sended comment!!');
                    }
                });
            }
        }
    }
</script>

<style scoped>

    /*LOADING ANIMATION */

    .blob {
        width: 2rem;
        height: 2rem;
        background: rgba(230,230,230,0.85);
        border-radius: 50%;
        position: absolute;
        left: calc(50% - 1rem);
        top: calc(50% - 1rem);
        box-shadow: 0 0 1rem rgba(255, 255, 255, 0.25);
    }

    .blob-2 { animation: animate-to-2 1.5s infinite; }
    .blob-3 { animation: animate-to-3 1.5s infinite; }
    .blob-1 { animation: animate-to-1 1.5s infinite; }
    .blob-4 { animation: animate-to-4 1.5s infinite; }
    .blob-0 { animation: animate-to-0 1.5s infinite; }
    .blob-5 { animation: animate-to-5 1.5s infinite; }

    @keyframes animate-to-2 {
        25%, 75% { transform: translateX(-1.5rem) scale(0.75); }
        95% { transform: translateX(0rem) scale(1); }
    }

    @keyframes animate-to-3 {
        25%, 75% { transform: translateX(1.5rem) scale(0.75); }
        95% { transform: translateX(0rem) scale(1); }
    }

    @keyframes animate-to-1 {
        25% { transform: translateX(-1.5rem) scale(0.75); }
        50%, 75% { transform: translateX(-4.5rem) scale(0.6); }
        95% { transform: translateX(0rem) scale(1); }
    }

    @keyframes animate-to-4 {
        25% { transform: translateX(1.5rem) scale(0.75); }
        50%, 75% { transform: translateX(4.5rem) scale(0.6); }
        95% { transform: translateX(0rem) scale(1); }
    }

    @keyframes animate-to-0 {
        25% { transform: translateX(-1.5rem) scale(0.75); }
        50% { transform: translateX(-4.5rem) scale(0.6); }
        75% { transform: translateX(-7.5rem) scale(0.5); }
        95% { transform: translateX(0rem) scale(1); }
    }

    @keyframes animate-to-5 {
        25% { transform: translateX(1.5rem) scale(0.75); }
        50% { transform: translateX(4.5rem) scale(0.6); }
        75% { transform: translateX(7.5rem) scale(0.5); }
        95% { transform: translateX(0rem) scale(1); }
    }

    .center-content {
        display: flex;
        flex-direction: row;
        background-color: #fff;
        margin: 50px 10%;
    }

    .book-image {
        height: 400px;
    }

    .text-content {
        text-align: left;
        padding: 20px;
    }

    .checked {
        color: #FFD700;
    }

    .rate-button {
        display: inline-block;
        background-color: #FF8C00;
        padding: 15px 20px;
        font-weight: 900;
        text-transform: uppercase;
        border-radius: 10px;
        color: #fff;
        margin-top: 5px;
        cursor: default;
    }

    .overall-rating {
        font-size: 18px;
        color: #FF8C00;
    }

    .max-rating {
        font-size: 24px;
        font-weight: bold;
    }

    .middle-box {
        display: flex;
        flex-direction: row;
        margin-top: 20px;
        justify-content: space-between;
    }

    .take-btn {
        margin-top: 10px;
        display: inline-block;
        padding: 10px;
        background: #f0f0f0;
        border-radius: 10px;
        cursor: default;
    }

    .comment-section {
        display: flex;
        flex-direction: row;
        margin-top: 20px;
    }

    .sendBtn {
        margin-left: 20px;
        background: #000080;
        padding: 10px;
        border-radius: 50%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .sendImg {
        height: 24px;
    }

    .comment-text {
        display: flex;
        align-items: center;
        width: 50%;
    }

    .comment-input {
        width: 100%;
        padding: 10px 5px;
        outline: none;
        border-radius: 5px;
    }
</style>