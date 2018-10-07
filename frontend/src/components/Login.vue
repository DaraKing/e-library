<template>
    <div id="Login">
        <img class="loginImage" src="../assets/books.svg">
        <div class="loginMessage">
            {{loginMessage}}
        </div>

        <div class="inputs">
            <input type="text" v-model="username" class="loginInput" placeholder="Your username"/>
            <div>
                <input type="password" v-model="password" class="loginInput" placeholder="Your password"/>
            </div>
            <input type="submit" class="submitBtn" value="Login" @click="login">
        </div>
        <div class="infoMessage">
            You dont have account.Please <router-link to="/register">register</router-link>
        </div>
    </div>
</template>

<script>
    import EventBus from '../event-bus';

    export default {
        name: "Login",
        data() {
            return {
                loginMessage: 'Welcome! Please login.',
                username: '',
                password: '',
            }
        },
        created: function() {
            if ( this.$cookie.get('rememberToken') ) {
                this.$router.push({name: 'home', query: {redirect: '/'} });
            }
        },
        methods: {
            login: function () {

                var loginData = {
                    "username" : this.username,
                    "password" : this.password,
                };

                this.$http.post('http://localhost:3030/login', loginData).then(response => {
                    if(response.status === 200) {
                        var rememberToken = response.bodyText;
                        this.$cookie.set('rememberToken', rememberToken, 365);
                        this.$router.push({name: 'home', query: {redirect: '/'} });
                        EventBus.$emit('logged');
                    }
                })

            }
        }
    }
</script>

<style scoped>
    @import url('https://fonts.googleapis.com/css?family=Lato');
    @import url('https://fonts.googleapis.com/css?family=Indie+Flower');

    * {
        font-family: 'Lato', sans-serif;
    }
    #Login {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        background: #fff;
        border-radius: 5px;
        padding: 40px 60px;
    }
    .loginImage {
        height: 150px;
    }
    .loginMessage {
        margin-top: 20px;
    }

    .inputs {
        margin-top: 25px;
    }

    .loginInput {
        font-family: 'Indie Flower', cursive;
        border: 0px;
        outline: none;
        margin-top: 10px;
        font-size: 20px;
    }

    .loginInput:focus {
        border-bottom: 1px solid #696969;
    }

    .submitBtn {
        margin-top: 25px;
        padding: 5px 10px;
        border: 1px solid #696969;
        background: transparent;
        border-radius: 7px;
        outline: none;
    }


</style>