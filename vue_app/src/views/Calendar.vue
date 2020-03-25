<template>
  <div class="calendar">
    <Week />
  </div>
</template>

<script>
  import { mapActions } from 'vuex'
  import Week from "@/components/Week.vue"


  export default {
    name: 'Calendar',
    components: {
      Week
    },
    created() {
      this.setMonthDays()
    },
    methods: {
      ...mapActions({
        setMonthDays: 'setMonthDays'
      }),
      setMonthDays() {
        this.setMonthDays()
      }
    }
  }
</script>