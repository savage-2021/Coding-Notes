import React from "react";
import fs from "fs";

const Post = ({ slug }) => <div>the slug for this page is: {slug}</div>;

export const getStaticPaths = async () => {
  const files = fs.readdirSync("posts");
  console.log(files);

  const paths = files.map(filename => ({
    params: {
      slug: filename.replace(".md", "")
    }
  }));
  console.log(paths);

  return {
    paths: paths,
    fallback: false
  };
};

export const getStaticProps = async ({ params: { slug } }) => {
  return {
    props: {
      slug
    }
  };
};

export default Post;
