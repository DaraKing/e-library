<template>
    <div id="header" v-show="showHeader">
        <div class="userInfo noselect" v-on:click="showDropDown">
            {{userInfo.username}}
            <img src="https://cdn2.iconfinder.com/data/icons/ios-7-icons/50/user_male2-512.png" class="userIcon">
            <div class="dropdown" v-show="dropdownShow">
                <div class="dropdown-content">Edit profile</div>
                <div class="dropdown-content" v-show="checkRole()"><router-link to="/books">All books</router-link></div>
                <div class="dropdown-content"><router-link to="/reserved-books">Reserved books</router-link></div>
                <div class="dropdown-content"> <router-link to="/logout">Logout</router-link></div>
            </div>
        </div>
    </div>
</template>

<script>
    import EventBus from '../event-bus';

    export default {
        name: "Header",
        data() {
            return {
                userInfo: [],
                dropdownShow: false,
                showHeader: false,
            }
        },
        created: function () {
            this.checkIsLoggedIn();
            this.getUserInfo();
            EventBus.$on('logged', () => {
                this.showHeader = true;
                this.getUserInfo();
            });

            EventBus.$on('logout', () => {
                this.showHeader = false;
            });

        },
        methods: {
            showDropDown: function(event) {
                this.dropdownShow = !this.dropdownShow;
            },
            checkIsLoggedIn: function () {
                if ( !(this.$cookie.get('rememberToken')) ) {
                    this.$router.push({name: 'login', query: {redirect: '/login'} });
                }else {
                    this.$nextTick(function () {
                        this.showHeader = true;
                    })
                }
            },
            getUserInfo: function() {
                if (this.$cookie.get('rememberToken') != null) {
                    var remember_token = {
                        "remember_token":  this.$cookie.get('rememberToken').trim().replace(/['"]+/g, ''),
                    };
                    this.$http.post('http://localhost:3030/userInfo', remember_token).then(response => {
                        this.userInfo = response.body;
                    });
                }
            },
            checkRole: function () {
                return this.userInfo.role_id == 1;
            }
        }
    }
</script>

<style scoped>
    #header {
        background: #fff;
        width: 100%;
        height: 64px;
    }

    .noselect {
        -webkit-touch-callout: none; /* iOS Safari */
        -webkit-user-select: none; /* Safari */
        -khtml-user-select: none; /* Konqueror HTML */
        -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
        user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
    }

    .userInfo {
        height: 100%;
        float: right;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 24px;
        font-weight: bold;
        cursor: default;
        width: 200px;
    }

    .userIcon {
        height: 52px;
        margin-left: 20px;
    }

    .dropdown {
        top: 64px;
        position: absolute;
        background: #fff;
        width: 200px;
    }

    .dropdown-content {
        padding: 10px 0px;
        font-size: 18px;
    }

    .dropdown-content > a {
        text-decoration: none;
        color: inherit;
    }

    .dropdown-content:hover {
        background: #696969;
        color: #fff;
    }
</style>