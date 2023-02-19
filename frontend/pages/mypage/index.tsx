import Segment from '../../components/Segment'
import Header from '../../components/Header'
import { Box, Grid, GridItem } from '@chakra-ui/react'

const MyPage = () => {
  return (
    <>
      <Header></Header>
      <Box height={'20em'} background={'gray.100'}></Box>
      <Segment>
        <Grid templateColumns='repeat(5, 1fr)' templateRows='repeat(2, 1fr)'>
          <GridItem colSpan={2}>ユーザー名</GridItem>
          <GridItem colSpan={3}>wallet address</GridItem>
          <GridItem colSpan={2}>自己紹介文</GridItem>
        </Grid>
      </Segment>
    </>
  )
}

export default MyPage
