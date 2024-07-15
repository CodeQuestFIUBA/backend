import { red } from '@mui/material/colors';
import { createTheme } from '@mui/material/styles';

// A custom theme for this app
const theme = createTheme({
    palette: {
        primary: {
            main: '#F6B75B',
        },
        secondary: {
            main: '#E84F3B',
        },
        error: {
            main: red.A400,
        },
    },
});

export default theme;
