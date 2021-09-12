import { MustBeOneOf } from "./Validator";

const validColors = [
  "",
  "primary",
  "info",
  "success",
  "warning",
  "danger",
  "dark",
  "light",
  "white",
  "black",
];

export default {
  computed: {
    fgColor: function() {
      return "is-" + this.color;
    },
    textColor: function() {
      let base = "has-text-";
      switch (this.color) {
        case "warning":
          return base + "dark";
        case "light":
          return base + "dark";
        case "white":
          return base + "dark";
      }

      return base + "white";
    },
    bgColor: function() {
      return "has-background-" + this.color;
    },
  },
  props: {
    color: {
      type: String,
      default: "light",
      validator: (value) => {
        return MustBeOneOf("color", value, validColors);
      },
    },
  },
};
