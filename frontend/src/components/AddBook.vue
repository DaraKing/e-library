<template>
    <div>
        <div>Add new book</div>
        <form enctype="multipart/form-data" action="http://localhost:3030/add-book" method="post">
            <div><input type="text" name="title" placeholder="Book title" /></div>
            <div><input type="text" name="author" placeholder="Book author" /></div>
            <div><input type="textarea" name="description" placeholder="Description about book"/></div>
            <div><input type="quantity" name="quantity" placeholder="Quantity"/></div>
            <div><input type="text" name="page_number" placeholder="Page number" /></div>
            <div><input type="file" name="uploadfile" /></div>
            <!--<input type="hidden" name="token" value=""/>-->
            <div><input type="submit" value="upload"/></div>
        </form>
    </div>
</template>

<script>
    export default {
        name: "AddBook",
        data() {
            return {
            }
        },
        methods: {

        }
    }
</script>

<style scoped>

</style>