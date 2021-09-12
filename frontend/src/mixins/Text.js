import { MustBeOneOf } from "./Validator";

const validSizes = [1, 2, 3, 4, 5, 6, 7];
const validWeights = ["bold", "regular", "light"];

export default {
  computed: {
    simpleWeight: function() {
      if (this.weight) {
        return "has-text-weight" + this.color;
      }
      return "";
    },
    simpleSize: function() {
      if (this.size) {
        return "is-size-" + this.size;
      }
      return "";
    },
  },
  props: {
    weight: {
      type: String,
      default: "regular",
      validator: function(value) {
        return MustBeOneOf("weight", value, validWeights);
      },
    },
    size: {
      type: Number,
      validator: function(value) {
        return MustBeOneOf("size", value, validSizes);
      },
    },
  },
};
