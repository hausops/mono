// material design 3 pallete

// export const neutral = {
//   50: '#fafafa',
//   100: '#f5f5f5',
//   200: '#eeeeee',
//   300: '#e0e0e0',
//   400: '#bdbdbd',
//   500: '#9e9e9e',
//   600: '#757575',
//   700: '#616161',
//   800: '#424242',
//   900: '#212121',
//   A100: '#f5f5f5',
//   A200: '#eeeeee',
//   A400: '#bdbdbd',
//   A700: '#616161',
// };

export const neutral = {
  0: '#000000',
  10: '#1d1a22',
  20: '#322f37',
  30: '#49454f',
  40: '#605d66',
  50: '#79747e',
  60: '#938f99',
  70: '#aea9b4',
  80: '#cac4d0',
  90: '#e7e0ec',
  95: '#f5eefa',
  99: '#fffbfe',
  100: '#ffffff',
};

// export const primary = '#002c7b';
// export const primary = '#00316b';
export const primary = '#3f51b5';
// export const secondary = '#fb728e';
// export const secondary = '#ffb703';
export const secondary = '#48cae4'; // monochrome

// export const primaryPallete = {
//   50: '#e0e6ed',
//   100: '#b3c1d3',
//   200: '#8098b5',
//   300: '#4d6f97',
//   400: '#265081',
//   500: '#00316b',
//   600: '#002c63',
//   700: '#002558',
//   800: '#001f4e',
//   900: '#00133c',
//   a100: '#728fff',
//   a200: '#3f66ff',
//   a400: '#0c3eff',
//   a700: '#0031f1',
// };

export const primaryPallete = {
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

// monochrome
export const secondaryPallete = {
  50: '#e9f9fc',
  100: '#c8eff7',
  200: '#a4e5f2',
  300: '#7fdaec',
  400: '#63d2e8',
  500: '#48cae4',
  600: '#41c5e1',
  700: '#38bddd',
  800: '#30b7d9',
  900: '#21abd1',
  a100: '#ffffff',
  a200: '#d7f6ff',
  a400: '#a4eaff',
  a700: '#8be4ff',
};

// --md-sys-color-primary: #4355b9;
// --md-sys-color-primary-rgb: 67 85 185;

// --md-sys-color-surface-1: rgb(var(--md-sys-color-primary-rgb) / 0.05);
// --md-sys-color-surface-2: rgb(var(--md-sys-color-primary-rgb) / 0.08);
// --md-sys-color-surface-3: rgb(var(--md-sys-color-primary-rgb) / 0.11);
// --md-sys-color-surface-4: rgb(var(--md-sys-color-primary-rgb) / 0.12);
// --md-sys-color-surface-5: rgb(var(--md-sys-color-primary-rgb) / 0.14);

const surface = {
  1: 'rgb(67 85 185 / 0.05)',
};

export const background = {
  $: '#f6f6f6',
  hovered: '#f1f1f1',
  pressed: '#ececec',
  selected: '#ececec',
};

export const text = {
  $: '#363636',
  disabled: '#8c9196',
  primary: {
    $: primaryPallete[35], // hsl(230, 52%, 44%)
    hovered: '#32449f', // hsl(230, 52%, 41%)
    pressed: '#2f3f93', // hsl(230, 52%, 38%)
  },
};

export const divider = '#e1e3e5';
