<template>
    <div>
        <!--<div id="header">-->
            <!--<div class="userInfo noselect" v-on:click="showDropDown">-->
                <!--{{userInfo.username}}-->
                <!--<img src="https://cdn2.iconfinder.com/data/icons/ios-7-icons/50/user_male2-512.png" class="userIcon">-->
                <!--<div class="dropdown" v-show="dropdownShow">-->
                    <!--<div class="dropdown-content">Edit profile</div>-->
                    <!--<div class="dropdown-content" v-show="checkRole()"><router-link to="/books">All books</router-link></div>-->
                    <!--<div class="dropdown-content">Reserved books</div>-->
                    <!--<div class="dropdown-content"> <router-link to="/logout">Logout</router-link></div>-->
                <!--</div>-->
            <!--</div>-->
        <!--</div>-->
    </div>
</template>

<script>
    export default {
        name: "Home",
        data() {
            return {
                //
            }
        },
        created: function () {
            this.checkIsLoggedIn();
        },
        methods: {
            checkIsLoggedIn: function () {
                if ( !(this.$cookie.get('rememberToken')) ) {
                    this.$router.push({name: 'login', query: {redirect: '/login'} });
                }
            },
        }
    }
</script>

<style scoped>

</style>