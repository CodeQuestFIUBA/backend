import {AuthLayout} from "../layout/AuthLayout";
import {Button, Grid, Link, TextField, Typography} from "@mui/material";
import {useForm} from "react-hook-form";
import {useAppDispatch, useAppSelector} from "../../hooks";
import {checkingAuthentication, checkingRegisterAuthentication} from "../../store/slice/auth/thunks";
import {Link as RouterLink} from "react-router-dom";
import {AuthForm} from "../../models/Forms";
import {useMemo, useState} from "react";
import {AuthStatus} from "../../store";

export const LoginPage = () => {

    const { status, errorMessage } = useAppSelector(state => state.auth);
    const dispatch = useAppDispatch();
    const { register, handleSubmit, formState: { errors } } = useForm<AuthForm>();
    const isAuthenticating = useMemo(() => status === AuthStatus.CHECKING, [status]);
    const [isLogin, setIsLogin] = useState(true);

    const onLogin = async (data: AuthForm) => {
        if (isLogin) {
            dispatch(checkingAuthentication(data));
        } else {
            dispatch(checkingRegisterAuthentication(data));
        }

    }

    return (
      <div>
          <AuthLayout title={isLogin ? "Ingresar" : "Registrarse"}>
              <form onSubmit={handleSubmit(onLogin)}>
                  <Grid container>

                      {
                          !isLogin && (
                              <Grid item xs={ 12 } sx={{ mt: 2, minHeight: "80px" }}>
                                  <TextField
                                    error={ !!errors.name }
                                    helperText={errors.name?.message }
                                    label="Nombre"
                                    type="text"
                                    placeholder="Ingresa tu nombre"
                                    fullWidth
                                    {
                                        ...register("name", {
                                            required: 'Debe ingresar su email',
                                        })
                                    }
                                  />
                              </Grid>
                          )
                      }

                      <Grid item xs={ 12 } sx={{ mt: 2, minHeight: "80px" }}>
                          <TextField
                            error={ !!errors.email }
                            helperText={errors.email?.message }
                            label="Email"
                            type="text"
                            placeholder="Ingresa tu email"
                            fullWidth
                            {
                                ...register("email", {
                                    required: 'Debe ingresar su email',
                                })
                            }
                          />
                      </Grid>

                      <Grid item xs={ 12 } sx={{ mt: 2 }}>
                          <TextField
                            error={ errors.password?.type === 'required' }
                            helperText={errors.password?.message }
                            label="Contaseña"
                            type="password"
                            placeholder="Ingresa tu contraseña"
                            fullWidth
                            {
                                ...register("password", {
                                    required: 'Debe ingresar su contraseña'
                                })
                            }
                          />
                      </Grid>


                      <Grid
                        container
                        spacing={ 2 }
                        sx={{ mb: 2, mt: 2 }}
                      >
                          <Grid item xs={12} >
                              <Button type="submit" variant="contained" fullWidth disabled={isAuthenticating} style={{color: 'white'}}>
                                  {isLogin ? 'Login' : 'Registrarse'}
                              </Button>
                          </Grid>
                      </Grid>

                      <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        alignItems="center"
                      >
                          {
                            !!errorMessage && <Typography variant="inherit" style={{color: "red"}}> Error al iniciar sesión</Typography>
                          }
                      </Grid>

                      <Grid
                        container
                        direction="row"
                        justifyContent="end"
                        alignItems="center"
                      >
                          <div style={{cursor: 'pointer'}} onClick={() => setIsLogin(!isLogin)}>
                              <Typography variant="inherit" style={{color: "red"}}>{isLogin ? 'Registrarse' : 'Ingresar'}</Typography>
                          </div>
                      </Grid>

                  </Grid>
              </form>
          </AuthLayout>

      </div>

    )
}
