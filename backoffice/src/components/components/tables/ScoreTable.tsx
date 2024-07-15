import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {useAppDispatch} from "../../../hooks";
import React, {useEffect, useState} from "react";
import {Backdrop, Button, CircularProgress, IconButton, Pagination} from "@mui/material";
import {useNavigate, useParams} from "react-router-dom";
import Visibility from "@mui/icons-material/Visibility";

import {Score} from "../../../models/Score";
import {getScoresByUserId} from "../../../store/slice/scores/thunks";

export const ScoreTable = () => {
    const navigate = useNavigate();
    let { id } = useParams();
    const dispatch = useAppDispatch();
    const [showLoading, setShowLoading] = useState(true);
    const [ scores, setScores ] = useState<Score[]>([]);

    const getScores = async (classRoomId: string) => {
        setShowLoading(true);
        // @ts-ignore
        const result = await dispatch(getScoresByUserId(classRoomId));
        // @ts-ignore
        setScores(result.scores || []);
        setShowLoading(false);
    }

    const openClassRoom = (id: string) => {
        navigate(`/students/${id}`);
    }

    useEffect(() => {
        if (id) {
            getScores(id);
        }
    }, [])

    return (
      <>
          <TableContainer component={Paper}>
              <Table aria-label="simple table">
                  <TableHead>
                      <TableRow>
                          <TableCell
                            align="left"
                            width="20%"
                          >
                              Nivel
                          </TableCell>
                          <TableCell
                            align="left"
                            width="20%"
                          >
                              Sub-Nivel
                          </TableCell>
                          <TableCell
                            align="left"
                            width="20%"
                          >
                              Puntos
                          </TableCell>
                          <TableCell
                            align="left"
                            width="20%"
                          >
                              Completado
                          </TableCell>
                          <TableCell
                            align="left"
                            width="20%"
                          >
                              Intentos
                          </TableCell>
                      </TableRow>
                  </TableHead>
                  <TableBody>
                      {scores.map((score) => (
                        <TableRow
                          key={score.id}
                          sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell
                              align="left"
                              width="20%"
                            >
                                {score.level_title}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="20%"
                            >
                                {score.sub_level_title}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="20%"
                            >
                                {score.points} Pts
                            </TableCell>
                            <TableCell
                              align="left"
                              width="20%"
                            >
                                {score.points > 0 ? "Si" : "No"}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="20%"
                            >
                                {score.attempts}
                            </TableCell>
                        </TableRow>
                      ))}
                  </TableBody>
              </Table>
          </TableContainer>

          <Backdrop
            sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
            open={showLoading}
          >
              <CircularProgress color="inherit" />
          </Backdrop>

      </>
    );
}

