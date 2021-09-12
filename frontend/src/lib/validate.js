/**
 * Validate a prop with a list of values
 * @param {string} prop - The prop title
 * @param {*} value - The value to check
 * @param {array} list - The list values
 * @returns {bool} Whether the field is valid
 */
export function MustBeOneOf(prop, value, list) {
  if (!list.includes(value)) {
    console.warn(`Prop "${prop}" got "${value}";\n\n Expected one of:`, list);
    return false;
  }
  return true;
}
