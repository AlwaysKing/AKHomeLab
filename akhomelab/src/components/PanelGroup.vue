<template>
  <div class="Group" :class="{ GroupTitlehover: this.hover }">
    <span
      class="GroupTitle"
      @click="HiddenOrShow"
      @mouseover="hover = true"
      @mouseout="hover = false"
    >
      {{ Title }}
    </span>
    <slot v-if="show"></slot>
    <div v-if="!show">
      <PanelItem @click="HiddenOrShow" :img="Image"></PanelItem>
    </div>
  </div>
</template>

<script>
// import { MDBContainer } from "mdb-vue-ui-kit";
import PanelItem from "./PanelItem.vue";
export default {
  name: "PanelGroup",
  components: { PanelItem },
  props: {
    Title: String,
    Image: String,
  },
  data() {
    return {
      show: true,
      hover: false,
    };
  },
  mounted() {
    this.show = localStorage.getItem("akhomelab." + this.Title) == "true";
  },
  methods: {
    HiddenOrShow() {
      this.show = !this.show;
      localStorage.setItem("akhomelab." + this.Title, this.show);
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.Group {
  margin: 10px 30px;
  border-radius: 6px;
}

.GroupTitle {
  margin-left: 20px;
  font-size: 25px;
  cursor: pointer;
}

.GroupTitlehover {
  background-color: rgba(119, 119, 119, 0.3);
}
</style>
