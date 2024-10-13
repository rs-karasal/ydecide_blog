import { useEffect } from "react";
import {
  InstagramOutlined,
  SendOutlined,
  LinkedinOutlined,
} from "@ant-design/icons";
import { Card } from "antd";
import coverPhoto from "../assets/images/contacts-cover-image.jpg";

const { Meta } = Card;

const ContactsPage = () => {
  
  useEffect(() => {
    document.title = "Contacts";
  }, []);

  const openLink = (url) => {
    window.open(url, "_blank");
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card
        style={{
          width: "100%",
          maxWidth: 600,
          minWidth: 300,
        }}
        cover={
          <img
            alt="profile picture"
            src={coverPhoto}
          />
        }
        actions={[
          <InstagramOutlined
            key="instagram"
            style={{ fontSize: "24px" }}
            onClick={() => openLink("https://www.instagram.com/rs_karasal/")}
          />,
          <SendOutlined
            key="telegram"
            style={{ fontSize: "24px" }}
            onClick={() => openLink("https://t.me/rs_karasal/")}
          />,
          <LinkedinOutlined
            key="linkedin"
            style={{ fontSize: "24px" }}
            onClick={() => openLink("https://www.linkedin.com/in/eres-karasal/")}
          />,
        ]}
      >
        <Meta
          title={<h1 className="text-lg sm:text-xl lg:text-2xl">Let’s be friends and keep in touch</h1>}
          description={
            <p className="text-sm sm:text-base lg:text-lg">
              “The people who are crazy enough to think they can change the
              world are the ones who do.”
            </p>
          }
        />
      </Card>
    </div>
  );
};

export default ContactsPage;