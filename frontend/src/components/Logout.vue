<template>
    
</template>

<script>
    import EventBus from '../event-bus';

    export default {
    name: "Logout",
    created: function () {
        this.removeCookie();
    },
    methods: {
        removeCookie: function() {
            this.$cookie.delete('rememberToken');
            this.$router.push({name: 'login', query: {redirect: '/login'} });
            EventBus.$emit('logout');
        },
    }
}
</script>

<style scoped>

</style>