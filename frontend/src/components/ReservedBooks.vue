<template>
    <div>
        <div class="container">
            <h2>Reserved books</h2>
            <table>
                <tr>
                    <th>Username</th>
                    <th>Book title</th>
                    <th>Expiration date</th>
                </tr>
                <tr v-for="reservedBook in reservedBooks">
                    <td>{{reservedBook.username}}</td>
                    <td>{{reservedBook.title}}</td>
                    <td>{{reservedBook.expiration_date}}</td>
                    <td>
                    <b-btn class="noButton" v-b-tooltip.hover title="Book is returned"  v-bind:id="reservedBook.id"v-on:click="setAsReturn"><font-awesome-icon :icon="['fas', 'undo']" /></b-btn>
                    </td>
                </tr>
            </table>
        </div>

    </div>
</template>

<script>
    export default {
        name: "ReservedBooks",
        data() {
            return {
                reservedBooks: [],
            }
        },
        created: function () {
            this.getAllReservedBooks();
        },
        methods: {
            getAllReservedBooks: function () {
                this.$http.get('http://localhost:3030/get-reserved-books').then(response => {
                    this.reservedBooks = response.body;
                    console.log(this.reservedBooks);
                })
            },
            setAsReturn: function() {
                var id = event.currentTarget.id;

                this.$http.delete('http://localhost:3030/set-as-return/'+id).then(response => {
                    if(response.status === 200) {
                        this.getAllReservedBooks();
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .container {
        margin: 50px 10% 0px 10%;
        background: #fff;
        padding: 20px;
    }

    table {
        margin-top: 20px;
        width: 100%;
    }

    tr > td {
        padding: 10px;
    }

    .noButton {
        border: 0px;
        outline: none;
        background: transparent;
    }
</style>