const primary = {
  0: '#000000',
  10: '#00105c',
  20: '#08218a',
  25: '#1b2f95',
  30: '#293ca0',
  35: '#3649ac',
  40: '#4355b9',
  50: '#5d6fd4',
  60: '#7789f0',
  70: '#97a5ff',
  80: '#bac3ff',
  90: '#dee0ff',
  95: '#f0efff',
  98: '#fbf8ff',
  99: '#fefbff',
  100: '#ffffff',
};

const neutral = {
  0: '#000000',
  10: '#121212',
  20: '#1f1f1f',
  25: '#252525',
  30: '#2c2c2c',
  35: '#333333',
  40: '#3a3a3a',
  50: '#494949',
  60: '#585858',
  70: '#787878',
  80: '#9d9d9d',
  90: '#cbcbcb',
  95: '#e8e8e8',
  98: '#f6f6f6',
  99: '#fbfbfb',
  100: '#ffffff',
};

const background = {
  $: neutral[98], // hsl(0, 0, 96%)
  hovered: '#f1f1f1', // hsl(0, 0, 95%)
  pressed: '#ececec', // hsl(0, 0, 93%)
  selected: '#ececec',
  transparent: 'transparent',
};

const surface = {
  $: neutral[100],
};

const text = {
  $: neutral[35],
  muted: neutral[70],
  // disabled: '#8c9196',
  primary: {
    $: primary[35], // hsl(230, 52%, 44%)
    // hovered: '#32449f', // hsl(230, 52%, 41%)
    // pressed: '#2f3f93', // hsl(230, 52%, 38%)
  },
};

const divider = neutral[95];

export const color = {
  primary,
  neutral,
  background,
  surface,
  text,
  divider,
};
