<template>
  <div id="message">
    <div class="col mb-5 mt-3">
      <div v-if="create"
        class="col mb-5 mt-3 mr-3 alert alert-success" role="alert">
        <button
          @click.prevent="closeAlert()"
          type="button"
          class="close"
          data-dismiss="alert" aria-label="Close">
        <span aria-hidden="true">&times;</span>
        </button>
        <h4 class="alert-heading">{{subject}} Created!</h4>
        <hr>
        <p>The {{subject}} was added Successfuly</p>
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
        <h4 class="alert-heading">Could not create a new {{subject}}!</h4>
        <hr>
        <p>The {{subject}} was not created</p>
      </div>
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