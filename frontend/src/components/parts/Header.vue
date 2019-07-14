<template>
  <div id="headerTitle">
    <div class="col mb-5 mt-3">
      <h1>{{ initTitle }}</h1>
    </div>
  </div>
</template>
<script>
export default {
  name: "headerTitle",
  props: ['initTitle'],
}
</script>
