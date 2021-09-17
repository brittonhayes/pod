import { MustBeOneOf } from "../lib/validate";

export enum Color {
  Empty = "",
  Primary = "primary",
  Success = "success",
  Warning = "warning",
  Danger = "danger",
  Info = "info",
  Dark = "dark",
  Light = "light",
  White = "white",
  Black = "black",
}

export enum Prefix {
  Is = "is-",
  HasBackground = "has-background-",
  HasText = "has-text-",
}

export function ColorPrefix(prefix: Prefix, color: Color): string {
  return prefix + color;
}

export function IsColor(color: Color): string {
  return ColorPrefix(Prefix.Is, color);
}

export function BackgroundColor(color: Color): string {
  return Prefix.HasBackground + color;
}

export function TextColor(color: Color): string {
  switch (color) {
    case Color.Warning:
      return Prefix.HasText + "dark";
    case Color.Light:
      return Prefix.HasText + "dark";
    case Color.Dark:
      return Prefix.HasText + "white";
    default:
      return "";
  }
}

export const Palette = [
  Color.Empty,
  Color.Primary,
  Color.Info,
  Color.Success,
  Color.Warning,
  Color.Danger,
  Color.Dark,
  Color.Light,
  Color.White,
  Color.Black,
];

export default {
  props: {
    color: {
      type: String,
      default: Color.White,
      validator: (value: Color) => {
        return MustBeOneOf("color", value, Palette);
      },
    },
  },
};
