import {ArtLayout} from "../layout/ArtLayout";
import {RoomTable} from "../components/tables/RoomTable";

export const ClassRoomPage = () => {

    return (
        <ArtLayout
            title={"Aulas"}
        >
            <RoomTable />
        </ArtLayout>
    )
}
