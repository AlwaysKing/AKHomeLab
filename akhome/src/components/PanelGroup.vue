<template>
  <div class="Group" :class="{ GroupTitlehover: this.hover, SmallIcon: SmallIcon }">
    <span class="GroupTitle" @click="HiddenOrShow" @mouseover="hover = true" @mouseout="hover = false">
      {{ Title }}
    </span>
    <slot v-if="show || SmallIcon"></slot>
    <div v-if="!show && !SmallIcon">
      <PanelItem @click="HiddenOrShow" :img="Image"></PanelItem>
    </div>
  </div>
</template>

<script>
import PanelItem from "./PanelItem.vue";
export default {
  name: "PanelGroup",
  components: { PanelItem },
  props: {
    Title: String,
    Image: String,
    SmallIcon: Boolean,
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

.SmallIcon {
  min-height: 310px;
  min-width: 220px;
}
</style>
