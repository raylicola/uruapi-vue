export const abbreviateText = (text, length, suffix) => {
  return text.substring(0, length) + (text.length > length ? suffix: '')
}