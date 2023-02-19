import Segment from '../../components/Segment'
import Header from '../../components/Header'
import {
  Box,
  Grid,
  GridItem,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Input,
} from '@chakra-ui/react'
import PurchasedArticleCard from '../../components/Card/PurchasedArticleCard'
import MyArticleCard from '../../components/Card/MyArticleCard'
import NFTCard from '../../components/Card/NFTCard'

const MyPage = () => {
  return (
    <>
      <Header></Header>
      <Box height={'20em'} background={'gray.100'}></Box>
      <Segment>
        <Box>ユーザー名</Box>
        <Box>wallet address</Box>
        <Box>自己紹介文</Box>
      </Segment>
      <Segment>
        <Tabs variant='enclosed'>
          <TabList>
            <Tab>購入した記事</Tab>
            <Tab>自分の記事</Tab>
            <Tab>持っている収益NFT</Tab>
          </TabList>

          <TabPanels>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <PurchasedArticleCard />
                </GridItem>
              </Grid>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <MyArticleCard />
                </GridItem>
              </Grid>
            </TabPanel>
            <TabPanel>
              <Input placeholder='検索' background={'gray.100'} />
              <Grid templateColumns='repeat(3, 1fr)'>
                <GridItem colSpan={1}>
                  <NFTCard />
                </GridItem>
              </Grid>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </Segment>
    </>
  )
}

export default MyPage
