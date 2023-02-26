import Segment from '../../components/Segment'
import Header from '../../components/Header'
import { Box, Tabs, TabList, TabPanels, Tab, TabPanel } from '@chakra-ui/react'
import MyArticleCardTab from '../../components/MyPageTabs/MyArticleCardTab'
import NFTCardTab from '../../components/MyPageTabs/NFTCardTab'
import PurchasedArticleTab from '../../components/MyPageTabs/PurchasedArticleTab'

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
              <PurchasedArticleTab />
            </TabPanel>
            <TabPanel>
              <MyArticleCardTab />
            </TabPanel>
            <TabPanel>
              <NFTCardTab />
            </TabPanel>
          </TabPanels>
        </Tabs>
      </Segment>
    </>
  )
}

export default MyPage
