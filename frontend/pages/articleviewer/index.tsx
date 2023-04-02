import { Grid, GridItem } from '@chakra-ui/react'
import Header from '../../components/Header'
import ArticleColumn from '../../components/ArticleViewer/ArticleColumn'
import SideColumn from '../../components/ArticleViewer/SideColumn'

const ArticleViewer = () => {
  return (
    <>
      <Header></Header>
      <Grid templateColumns='repeat(4, 1fr)'>
        <GridItem colSpan={3}>
          <ArticleColumn />
        </GridItem>
        <GridItem colSpan={1}>
          <SideColumn />
        </GridItem>
      </Grid>
    </>
  )
}

export default ArticleViewer
