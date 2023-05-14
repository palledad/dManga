import Segment from '../../components/Segment'
import Header from '../../components/Header'
import { Grid, GridItem, Input } from '@chakra-ui/react'
import SearchArticleCard from '../../components/Card/SearchArticleCard'
import Link from 'next/link'
import { GetStaticProps } from 'next'
import { Article, DefaultApi } from '../../api'

export const getStaticProps: GetStaticProps = async () => {
  const articleAlias = ['2d84s4h0a']

  const promises = articleAlias.map((alias) =>
    new DefaultApi().getArticle(alias).then((x) => x.data),
  )

  const ret = await Promise.all<Promise<Article>[]>(promises)

  return {
    props: {
      articlesData: ret,
    },
  }
}

type MyPageProps = {
  articlesData: Article[]
}
const MyPage = ({ articlesData }: MyPageProps) => {
  return (
    <>
      <Header></Header>
      <Segment>
        <Input placeholder='検索' background={'gray.100'} />
        <Grid templateColumns='repeat(3, 1fr)'>
          {articlesData.map((x) => (
            <GridItem key={x.alias} colSpan={1}>
              <Link href={`articleviewer/${x.alias}`}>
                <SearchArticleCard
                  key={x.alias}
                  title={x.title}
                  author={x.author_address}
                  createdAt={x.created_at}
                  updatedAt={x.updated_at}
                />
              </Link>
            </GridItem>
          ))}
        </Grid>
      </Segment>
    </>
  )
}

export default MyPage
