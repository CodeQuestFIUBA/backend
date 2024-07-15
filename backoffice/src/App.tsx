import {AppRouter} from "./router/AppRouter";
import {AppTheme} from "./theme/AppTheme";
import 'react-toastify/dist/ReactToastify.css';

export const App = () => {
  return (
      <AppTheme>
        <AppRouter />
      </AppTheme>
  )
}
