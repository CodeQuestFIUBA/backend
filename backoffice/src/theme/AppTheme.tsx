import { CssBaseline, ThemeProvider } from '@mui/material';
import theme from './theme';

export const AppTheme = ({children}: any) => {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            { children }
        </ThemeProvider>
    )
}
