<template>
  <div class="col mb-3 mt-3" id="message">
    <div v-if="create"
      class="col mb-5 mt-3 mr-3 alert alert-success" role="alert">
      <button
        @click.prevent="closeAlert()"
        type="button"
        class="close"
        data-dismiss="alert" aria-label="Close">
      <span aria-hidden="true">&times;</span>
      </button>
      <h4 class="alert-heading">{{subject}}</h4>
      <hr>
      <p>The {{subject}} Successfuly</p>
    </div>

    <div v-if="!create && errorCreate != null"
      class="col mb-5 mt-3 mr-3 alert alert-warning"
      role="alert">

      <button
        @click.prevent="closeAlert()"
        type="button"
        class="close"
        data-dismiss="alert"
        aria-label="Close">
        <span aria-hidden="true">&times;</span>
      </button>
      <h4 class="alert-heading">Could not {{subject}}!</h4>
      <hr>
      <p>{{subject}} Failed</p>
    </div>
  </div>
</template>
<script>
export default {
  name: "message",
  props: {
    create:  Boolean,
    errorCreate: null,
    subject: String,
  },
  data: () => ({
  }),
  methods: {
    closeAlert:function(){
      this.$emit('close-alert');
    }
  }
}
</script>
<style lang="scss" scoped>
</style>