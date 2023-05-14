import { Grid, GridItem } from '@chakra-ui/react'
import ArticleColumn from '../../components/ArticleViewer/ArticleColumn'
import SideColumn from '../../components/ArticleViewer/SideColumn'
import Header from '../../components/Header'
import { GetServerSideProps } from 'next'
import { Article, DefaultApi } from '../../api'

export const getServerSideProps: GetServerSideProps = async (context) => {
  const articleId = context.params?.article_id
  if (!articleId) {
    return {
      props: {
        articleData: {},
      },
    }
  }
  const dataPromise = await new DefaultApi().getArticle(String(articleId))
  const data = dataPromise.data
  return {
    props: {
      articleData: data,
    },
  }
}

type MyPageProps = {
  articleData: Article
}

const ArticleViewer = ({ articleData }: MyPageProps) => {
  return (
    <>
      <Header></Header>
      <Grid templateColumns='repeat(4, 1fr)'>
        {/*<p>{JSON.stringify(articleData)}</p>*/}
        <GridItem colSpan={3}>
          <ArticleColumn articleJson={articleData} />
        </GridItem>
        <GridItem colSpan={1}>
          <SideColumn />
        </GridItem>
      </Grid>
    </>
  )
}

export default ArticleViewer
