<template>
  <div class="">
    <p>{{ name }}</p>
    <p>{{ email }}</p>
    <p>{{ description }}</p>
  </div>
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      required: true,
    },
    email: {
      type: String,
    },
    description: {
      type: String,
    },
  },
};
</script>
