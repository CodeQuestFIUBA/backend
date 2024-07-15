import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {useAppDispatch} from "../../../hooks";
import {getClassRoom} from "../../../store/slice/classRoom/thunks";
import React, {useEffect, useState} from "react";
import {ClassRoom} from "../../../models/Art";
import {Backdrop, CircularProgress, IconButton, Pagination} from "@mui/material";
import Visibility from '@mui/icons-material/Visibility';
import {useNavigate} from "react-router-dom";
import FileCopyIcon from '@mui/icons-material/FileCopy';

export const RoomTable = () => {
    const navigate = useNavigate();
    const dispatch = useAppDispatch();
    const [showLoading, setShowLoading] = useState(true);
    const [ classRooms, setClassRooms ] = useState<ClassRoom[]>([]);
    const [ paginationProps, setPaginationProps ] = useState<{page: number, limit: number, total: number}>({
        page: 0,
        limit: 10,
        total: 0
    });

    const getAllClassRoom = async (page: number) => {
        setShowLoading(true);
        const result = await dispatch(getClassRoom(page, paginationProps.limit));
        setClassRooms(result.classRoom || []);
        setPaginationProps({
            ...paginationProps,
            page: page,
            total: result.total
        });
        setShowLoading(false);
    }

    const handleChange = async (event: React.ChangeEvent<unknown>, value: number) => {
        if (value - 1 !== paginationProps.page) {
            await getAllClassRoom(value - 1);
        }
    };

    const openClassRoom = (id: string) => {
        navigate(`/students/${id}`);
    }

    useEffect(() => {
        getAllClassRoom(paginationProps.page);
    }, [])

    return (
        <>
            <TableContainer component={Paper}>
                <Table aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell
                                align="left"
                                width="25%"
                            >
                                Aula
                            </TableCell>
                            <TableCell
                              align="left"
                              width="25%"
                            >
                                Código
                            </TableCell>
                            <TableCell
                              align="left"
                              width="25%"
                            >
                                Total de niveles completados
                            </TableCell>
                            <TableCell
                                align="left"
                                width="25%"
                            >
                                Ver aula
                            </TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {classRooms.map((classRoom) => (
                            <TableRow
                                key={classRoom.id}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell
                                    align="left"
                                    width="25%"
                                >
                                    {classRoom.name}
                                </TableCell>
                                <TableCell
                                  align="left"
                                  width="25%"
                                >
                                    <div style={{display: 'flex', alignItems: 'center'}}>
                                        <IconButton
                                          aria-label="delete"
                                          color="primary"
                                          onClick={() => {
                                              navigator.clipboard.writeText(classRoom.code).then(function() {
                                                  console.log('Text copied to clipboard successfully!');
                                              }, function(err) {
                                                  console.error('Could not copy text: ', err);
                                              });

                                          }}
                                        >
                                            <FileCopyIcon />
                                        </IconButton>
                                        {classRoom.code}
                                    </div>

                                </TableCell>
                                <TableCell
                                  align="left"
                                  width="25%"
                                >
                                    {classRoom.totalScores} / {classRoom.totalQuestions}
                                </TableCell>
                                <TableCell
                                    align="left"
                                    width="25%"
                                >
                                    <IconButton
                                        aria-label="delete"
                                        color="primary"
                                        onClick={() => { openClassRoom(classRoom.id) }}
                                    >
                                        <Visibility />
                                    </IconButton>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            {
                paginationProps.total > paginationProps.limit && (
                    <div style={{display: "flex", justifyContent: "flex-end", margin: "10px"}}>
                        <Pagination
                            count={paginationProps.total / paginationProps.limit}
                            page={paginationProps.page + 1}
                            onChange={handleChange}
                            color="primary"
                        />
                    </div>
                )
            }


            <Backdrop
                sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
                open={showLoading}
            >
                <CircularProgress color="inherit" />
            </Backdrop>

        </>
    );
}
