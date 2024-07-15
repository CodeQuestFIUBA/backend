import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {useAppDispatch} from "../../../hooks";
import React, {useEffect, useState} from "react";
import {Backdrop, CircularProgress, IconButton} from "@mui/material";
import {useNavigate, useParams} from "react-router-dom";
import Visibility from "@mui/icons-material/Visibility";
import {getStudentsByClassRoom} from "../../../store/slice/students/thunks";
import {Student} from "../../../models/Student";

export const StudentsTable = () => {
    const navigate = useNavigate();
    let { id } = useParams();
    const dispatch = useAppDispatch();
    const [showLoading, setShowLoading] = useState(true);
    const [ students, setStudents ] = useState<Student[]>([]);

    const getAllClassRoom = async (classRoomId: string) => {
        setShowLoading(true);
        const result = await dispatch(getStudentsByClassRoom(classRoomId));
        setStudents(result.users || []);
        setShowLoading(false);
    }

    const openScores = (id: string) => {
        navigate(`/score/${id}`);
    }

    useEffect(() => {
        if (id) {
            getAllClassRoom(id);
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
                            width="16%"
                          >
                              Nombre
                          </TableCell>
                          <TableCell
                            align="left"
                            width="16%"
                          >
                              Username
                          </TableCell>
                          <TableCell
                            align="left"
                            width="16%"
                          >
                              Email
                          </TableCell>
                          <TableCell
                            align="left"
                            width="16%"
                          >
                              Puntos
                          </TableCell>
                          <TableCell
                            align="left"
                            width="16%"
                          >
                              Niveles
                          </TableCell>
                          <TableCell
                            align="left"
                            width="16%"
                          >
                              Ver puntos
                          </TableCell>
                      </TableRow>
                  </TableHead>
                  <TableBody>
                      {students.map((student) => (
                        <TableRow
                          key={student.ID}
                          sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                {student.first_name} {student.last_name}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                {student.username}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                {student.email}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                {student.score} Pts
                            </TableCell>
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                {student.completed_levels} / {student.total_levels}
                            </TableCell>
                            <TableCell
                              align="left"
                              width="16%"
                            >
                                <IconButton
                                  aria-label="delete"
                                  color="primary"
                                  onClick={() => { openScores(student.ID) }}
                                >
                                    <Visibility />
                                </IconButton>
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

