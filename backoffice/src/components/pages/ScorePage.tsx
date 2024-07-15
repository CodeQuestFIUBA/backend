import {ArtLayout} from "../layout/ArtLayout";
import {StudentsTable} from "../components/tables/StudentsTable";
import {ScoreTable} from "../components/tables/ScoreTable";

export const ScorePage = () => {
    return (
        <ArtLayout
            title={"Puntos"}
            hideNew
            showBack
        >
            <ScoreTable />
        </ArtLayout>
    )
}
