import {ArtLayout} from "../layout/ArtLayout";
import {StudentsTable} from "../components/tables/StudentsTable";

export const StudentPage = () => {
    return (
        <ArtLayout
            title={"Alumnos"}
            hideNew
            showBack
        >
            <StudentsTable />
        </ArtLayout>
    )
}
