<template>
  <div class="container">
    <div @blur="test()">
      <vue-drag-resize
        @activated="activate()"
        @deactivated="deactivate()"
        :w="200"
        :h="200"
        v-on:resizing="resize"
        v-on:dragging="resize"
      >
        <div>
          <div v-if="!editing">
            <div v-if="active">
              <v-btn class="edit-button" icon @click="editing = !editing">
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
            </div>
            <v-md-preview :text="text"></v-md-preview>
          </div>
          <div v-else>
            <v-btn
              class="close-button"
              icon
              @click="editing = false"
              color="#F44336"
              ><v-icon>mdi-close</v-icon></v-btn
            >
            <v-md-editor v-model="text" height="400px"></v-md-editor>
          </div>
        </div>
      </vue-drag-resize>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      editing: false,
      active: false,
      width: 0,
      height: 0,
      top: 0,
      left: 0,
      text: '',
    }
  },
  methods: {
    resize(newRect) {
      this.width = newRect.width
      this.height = newRect.height
      this.top = newRect.top
      this.left = newRect.left
    },
    activate() {
      this.active = true
    },
    deactivate() {
      this.active = false
    },
  },
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  font-size: 15px;
}
.edit-button {
  position: absolute;
  top: 0%;
  right: 0%;
}
.close-button {
  position: absolute;
  top: -40px;
  right: -10px;
}
</style>
