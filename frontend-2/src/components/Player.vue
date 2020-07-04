<template>
  <div>
      <h2>{{name}}</h2>
      <h2>{{id}}</h2>
  </div>
</template>

<script>
export default{
    name: 'Player',
    props: {
        name: String,
        id: String
    }
}
</script>