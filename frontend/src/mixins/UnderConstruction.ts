import { SnackbarProgrammatic as Snackbar } from "buefy";

export const Notify = (message: string) => {
  Snackbar.open({
    type: "is-danger",
    position: "is-bottom-right",
    message: message,
  });
};

export default {
  mounted() {
    Notify("This page is under construction");
  },
};
