<template>
  <div class="week">
    <div v-for="day in dayArray" :key="day">
      <Day />
    </div>
  </div>
</template>

<script>
  import Day from '@/components/Day.vue'

  export default {
    name: 'Week',
    components: {
      Day
    },
    data() {
      return{
        // dayArray: [1,2,3]
        dayArray: this.$store.state.monthDays
      }
    }
  }
</script>